<script setup>
import { ref } from "vue";
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";
import { gql } from "@apollo/client/core";

const { $apollo } = useNuxtApp();

const loading = ref(false);

const SIGNUP_MUTATION = gql`
  mutation Signup($user: SignupInput!) {
    signup(user: $user) {
      message
      user {
        id
        name
        email
        role
      }
    }
  }
`;

const schema = yup.object({
  name: yup
    .string()
    .required("Full name is required")
    .min(3, "Name must be at least 3 characters"),

  email: yup.string().email("Invalid email").required("Email is required"),

  password: yup
    .string()
    .required("Password is required")
    .min(6, "Password must be at least 6 characters"),

  confirmPassword: yup
    .string()
    .required("Confirm your password")
    .oneOf([yup.ref("password")], "Passwords do not match"),
});

const onSubmit = async (values, { resetForm }) => {
  console.log("SUBMIT FIRED");
  loading.value = true;

  try {
    const { data } = await $apollo.mutate({
      mutation: SIGNUP_MUTATION,
      variables: {
        user: {
          name: values.name,
          email: values.email,
          password: values.password,
        },
      },
    });

    console.log("Registered User:", data.signup.user);

    resetForm();

    await navigateTo("/auth/login");
  } catch (err) {
    console.error(err);

    const message =
      err?.graphQLErrors?.[0]?.message ||
      err?.networkError?.result?.errors?.[0]?.message ||
      err?.message ||
      "Signup failed";

    alert(message);
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <div class="w-full max-w-md rounded-xl bg-zinc-900 p-8 shadow-lg">
    <h1 class="mb-8 text-center text-3xl font-bold text-white">
      Create Account
    </h1>

    <Form :validation-schema="schema" @submit="onSubmit" class="space-y-5">
      <div>
        <Field
          name="name"
          type="text"
          placeholder="Full Name"
          class="w-full rounded-lg border border-zinc-700 bg-zinc-800 px-4 py-3 text-white outline-none focus:border-lime-500"
        />
        <ErrorMessage name="name" class="mt-1 block text-sm text-red-500" />
      </div>

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

      <div>
        <Field
          name="confirmPassword"
          type="password"
          placeholder="Confirm Password"
          class="w-full rounded-lg border border-zinc-700 bg-zinc-800 px-4 py-3 text-white outline-none focus:border-lime-500"
        />
        <ErrorMessage
          name="confirmPassword"
          class="mt-1 block text-sm text-red-500"
        />
      </div>

      <button
        type="submit"
        :disabled="loading"
        class="w-full rounded-lg cursor-pointer bg-lime-500 py-3 font-semibold text-black transition hover:bg-lime-400 disabled:cursor-not-allowed disabled:opacity-50"
      >
        {{ loading ? "Creating Account..." : "Create Account" }}
      </button>
    </Form>

    <p class="mt-6 text-center text-sm text-zinc-400">
      Already have an account?
      <NuxtLink
        to="/auth/login"
        class="font-semibold text-lime-500 hover:text-lime-400"
      >
        Log in
      </NuxtLink>
    </p>
  </div>
</template>
