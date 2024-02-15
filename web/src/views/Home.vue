<script lang="ts" setup>
import { API_CONSTANTS } from "../constants/index";

const handleGithubLogin = () => {
  window.location.href = `${API_CONSTANTS.URL}/github`;
};
</script>

<template>
  <div class="flex flex-1 w-full h-full">
    <div class="card bg-base-200">
      <div class="card-body p-4">
        <button @click="handleGithubLogin()">Github</button>
      </div>
    </div>
  </div>
</template>
