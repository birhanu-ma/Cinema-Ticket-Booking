<script setup>
import { ref } from "vue";
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";
import { gql } from "@apollo/client/core";

const { $apollo } = useNuxtApp();
const { login } = useAuth();

const loading = ref(false);

const LOGIN_MUTATION = gql`
  mutation Login($user: LoginInput!) {
    login(user: $user) {
      message
      token
      type
      expires_in
    }
  }
`;

const schema = yup.object({
  email: yup
    .string()
    .email("Please enter a valid email")
    .required("Email is required"),

  password: yup
    .string()
    .required("Password is required")
    .min(6, "Password must be at least 6 characters"),
});

const decodeJwtRole = (token) => {
  try {
    const payloadBase64 = token.split(".")[1];

    const payloadJson = atob(
      payloadBase64.replace(/-/g, "+").replace(/_/g, "/"),
    );

    const payload = JSON.parse(payloadJson);

    return (
      payload?.["https://hasura.io/jwt/claims"]?.["x-hasura-default-role"] ||
      payload?.role ||
      null
    );
  } catch (err) {
    console.error("Failed to decode JWT:", err);
    return null;
  }
};

const onSubmit = async (values, { resetForm }) => {
  loading.value = true;

  try {
    const { data } = await $apollo.mutate({
      mutation: LOGIN_MUTATION,
      variables: {
        user: {
          email: values.email,
          password: values.password,
        },
      },
    });

    const token = data.login.token;

  
    login(token);

    resetForm();

    const role = decodeJwtRole(token);

    if (role === "admin") {
      await navigateTo("/admin/movies");
    } else {
      await navigateTo("/user/movies");
    }
  } catch (err) {
    console.error(err);

    const message =
      err?.graphQLErrors?.[0]?.message ||
      err?.networkError?.result?.errors?.[0]?.message ||
      err?.message ||
      "Login failed";

    alert(message);
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <div class="w-full max-w-md rounded-xl bg-zinc-900 p-8 shadow-lg">
    <h1 class="mb-8 text-center text-3xl font-bold text-white">Welcome Back</h1>

    <Form :validation-schema="schema" @submit="onSubmit" class="space-y-5">
      <div>
        <Field
          name="email"
          type="email"
          placeholder="Email Address"
          class="w-full rounded-lg border border-zinc-700 bg-zinc-800 px-4 py-3 text-white outline-none focus:border-lime-500"
        />
        <ErrorMessage name="email" class="mt-1 block text-sm text-red-500" />
      </div>

      <div>
        <Field
          name="password"
          type="password"
          placeholder="Password"
          class="w-full rounded-lg border border-zinc-700 bg-zinc-800 px-4 py-3 text-white outline-none focus:border-lime-500"
        />
        <ErrorMessage name="password" class="mt-1 block text-sm text-red-500" />
      </div>

      <button
        type="submit"
        :disabled="loading"
        class="w-full rounded-lg cursor-pointer bg-lime-500 py-3 font-semibold text-black transition hover:bg-lime-400 disabled:cursor-not-allowed disabled:opacity-50"
      >
        {{ loading ? "Signing In..." : "Sign In" }}
      </button>
    </Form>
  </div>
</template>
