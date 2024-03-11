import { createApp } from "vue";
import { OhVueIcon, addIcons } from "oh-vue-icons";
import { CoGoogle, CoGithub } from "oh-vue-icons/icons";

import "./style.css";
import router from "./router/routes";
import App from "./App.vue";

addIcons(CoGoogle, CoGithub);

const app = createApp(App);
app.use(router);
app.component("v-icon", OhVueIcon);
app.mount("#app");
