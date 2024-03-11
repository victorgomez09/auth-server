<script lang="ts">
import { reactive } from "vue";
import { useVuelidate } from "@vuelidate/core";
import { required, email } from "@vuelidate/validators";

import { API_CONSTANTS } from "../constants/index";

export default {
  setup() {
    const state = reactive({
      firstName: "",
      lastName: "",
      email: "",
      password: "",
    });
    const rules = {
      firstName: { required }, // Matches state.firstName
      lastName: { required }, // Matches state.lastName
      email: { required, email }, // Matches state.email
      password: { required },
    };

    const v$ = useVuelidate(rules, state);

    const handleGithubLogin = () => {
      window.location.href = `${API_CONSTANTS.URL}/github/login`;
    };

    return { state, v$, handleGithubLogin };
  },
  methods: {
    async submitForm() {
      const isFormCorrect = await this.v$.$validate();
      // you can show some extra alert to the user or just leave the each field to show it's `$errors`.
      if (!isFormCorrect) return;
      // actually submit form
    },
  },
};
</script>

<template>
  <div class="hero min-h-screen bg-base-200 w-full h-full">
    <div class="hero-content lg:flex-row-reverse w-full h-full">
      <div class="card shadow-md bg-base-100 w-6/12 h-full">
        <div class="card-body p-4 w-full h-full">
          <form class="p-2 overflow-auto" @submit="submitForm">
            <h2 class="card-title justify-center">Register</h2>
            <div class="form-control">
              <label class="label">
                <span class="label-text">First name</span>
              </label>
              <input
                type="text"
                placeholder="John"
                class="input input-bordered"
                v-model="state.firstName"
                :class="{ 'input-error': v$.firstName.$error }"
                @blur="v$.firstName.$touch"
              />
              <div class="label" v-if="v$.firstName.$error">
                <span class="label-text-alt text-error">{{
                  v$.firstName.$errors[0].$message
                }}</span>
              </div>
            </div>
            <!-- Last name -->
            <div class="form-control">
              <label class="label">
                <span class="label-text">Last name</span>
              </label>
              <input
                type="text"
                placeholder="Doe"
                class="input input-bordered"
                v-model="state.lastName"
                :class="{ 'input-error': v$.lastName.$error }"
                @blur="v$.lastName.$touch"
              />
              <div class="label" v-if="v$.lastName.$error">
                <span class="label-text-alt text-error">{{
                  v$.lastName.$errors[0].$message
                }}</span>
              </div>
            </div>
            <!-- Email -->
            <div class="form-control">
              <label class="label">
                <span class="label-text">Email</span>
              </label>
              <input
                type="email"
                placeholder="johndoe@gmail.com"
                class="input input-bordered"
                v-model="state.email"
                :class="{ 'input-error': v$.email.$error }"
                @blur="v$.email.$touch"
              />
              <div class="label" v-if="v$.email.$error">
                <span
                  class="label-text-alt text-error"
                  v-for="error in v$.email.$errors"
                >
                  {{ error.$message }}
                </span>
              </div>
            </div>
            <div class="form-control">
              <label class="label">
                <span class="label-text">Password</span>
              </label>
              <input
                type="password"
                placeholder="*********"
                class="input input-bordered"
                v-model="state.password"
                :class="{ 'input-error': v$.password.$error }"
                @blur="v$.password.$touch"
              />
              <div class="label" v-if="v$.password.$error">
                <span class="label-text-alt text-error">
                  {{ v$.password.$errors[0].$message }}
                </span>
              </div>
            </div>
            <div class="form-control mt-6">
              <button
                class="btn btn-primary"
                :disabled="!v$.$anyDirty || v$.$error"
              >
                Register
              </button>

              <span class="text-center font-thin mt-2">
                Already have an account?
                <router-link to="/login" class="text-info link ms-1">
                  Click here
                </router-link>
              </span>
            </div>
          </form>
          <div class="divider">OR</div>

          <div class="flex flex-col gap-2">
            <button class="btn">
              <v-icon name="co-google" />
              Login with Google
            </button>

            <button class="btn" @click="handleGithubLogin">
              <v-icon name="co-github" />
              Login with Github
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
