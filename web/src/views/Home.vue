<script lang="ts" setup>
import { API_CONSTANTS } from "../constants/index";

const handleGithubLogin = () => {
  window.location.href = `${API_CONSTANTS.URL}/github`;
};
</script>

<template>
  <div class="hero min-h-screen bg-base-200">
    <div class="hero-content w-full lg:flex-row-reverse">
      <div class="card w-6/12 shadow-lg bg-base-100">
        <div class="card-body">
          <form class="">
            <h2 class="card-title justify-center">Login</h2>
            <div class="form-control">
              <label class="label">
                <span class="label-text">Email</span>
              </label>
              <input
                type="email"
                placeholder="email"
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
                placeholder="password"
                class="input input-bordered"
                required
              />
              <label class="label">
                <a href="#" class="label-text-alt link link-hover"
                  >Forgot password?</a
                >
              </label>
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
