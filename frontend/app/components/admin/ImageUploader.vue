<script setup>
import { ref } from "vue";
import { gql } from "@apollo/client/core";

const props = defineProps({
  modelValue: {
    type: String,
    default: "",
  },
  label: {
    type: String,
    default: "Upload Image",
  },
});

const emit = defineEmits(["update:modelValue"]);

const { $apollo } = useNuxtApp();

const uploading = ref(false);
const errorMsg = ref("");
const fileInput = ref(null);

const UPLOAD_IMAGE = gql`
  mutation UploadImage($base64: String!, $filename: String!, $content_type: String!) {
    uploadImage(base64: $base64, filename: $filename, content_type: $content_type) {
      success
      message
      url
    }
  }
`;

const readFileAsBase64 = (file) => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();

    reader.onload = () => resolve(reader.result.split(",")[1]);
    reader.onerror = () => reject(new Error("Failed to read file"));

    reader.readAsDataURL(file);
  });
};

const handleFileChange = async (event) => {
  const file = event.target.files?.[0];

  if (!file) return;

  errorMsg.value = "";
  uploading.value = true;

  try {
    const base64 = await readFileAsBase64(file);

    const { data } = await $apollo.mutate({
      mutation: UPLOAD_IMAGE,
      variables: {
        base64,
        filename: file.name,
        content_type: file.type,
      },
    });

    if (!data.uploadImage.success) {
      errorMsg.value = data.uploadImage.message;
      return;
    }

    emit("update:modelValue", data.uploadImage.url);
  } catch (err) {
    console.error("Upload failed:", err);

    errorMsg.value =
      err?.graphQLErrors?.[0]?.message || "Upload failed, please try again";
  } finally {
    uploading.value = false;

    if (fileInput.value) {
      fileInput.value.value = "";
    }
  }
};

const triggerFileSelect = () => {
  fileInput.value?.click();
};
</script>

<template>
  <div class="flex flex-col gap-2">
    <label class="text-xs font-bold text-gray-400 uppercase tracking-wider">
      {{ label }}
    </label>

    <div
      class="relative w-full aspect-video bg-gray-900 border-2 border-dashed border-gray-800 rounded-xl overflow-hidden flex items-center justify-center cursor-pointer hover:border-lime-400 transition-colors"
      @click="triggerFileSelect"
    >
      <img
        v-if="modelValue"
        :src="modelValue"
        class="w-full h-full object-cover"
      />

      <div v-else-if="uploading" class="text-gray-500 text-sm">
        Uploading...
      </div>

      <div v-else class="text-gray-500 text-sm text-center px-4">
        Click to upload<br />
        <span class="text-xs">JPEG, PNG, or WebP — max 5MB</span>
      </div>

      <input
        ref="fileInput"
        type="file"
        accept="image/jpeg,image/png,image/webp"
        class="hidden"
        @change="handleFileChange"
      />
    </div>

    <p v-if="errorMsg" class="text-red-400 text-xs">
      {{ errorMsg }}
    </p>
  </div>
</template>