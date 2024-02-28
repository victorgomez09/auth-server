<script setup lang="ts">
import { API_CONSTANTS } from "../constants/index";

const handleGithubLogin = () => {
  window.location.href = `${API_CONSTANTS.URL}/github/login`;
};
</script>

<template>
  <div class="hero min-h-screen bg-base-200 w-full h-full">
    <div class="hero-content lg:flex-row-reverse w-full h-full">
      <div class="card shadow-md bg-base-100 w-6/12 h-full">
        <div class="card-body w-full h-full max-h-0">
          <form class="overflow-auto">
            <h2 class="card-title justify-center">Register</h2>
            <div class="form-control">
              <label class="label">
                <span class="label-text">Name</span>
              </label>
              <input
                type="text"
                placeholder="John Doe"
                class="input input-bordered"
                required
              />
            </div>
            <div class="form-control">
              <label class="label">
                <span class="label-text">Email</span>
              </label>
              <input
                type="email"
                placeholder="johndoe@gmail.com"
                class="input input-bordered"
                required
              />
            </div>
            <div class="form-control">
              <label class="label">
                <span class="label-text">Password</span>
              </label>
              <input
                type="password"
                placeholder="*********"
                class="input input-bordered"
                required
              />
            </div>
            <div class="form-control mt-6">
              <button class="btn btn-primary">Login</button>
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
