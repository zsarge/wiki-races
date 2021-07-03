<template>
  <div class="game">
    This is the view of the page:
    <br />
    dev: {{ dev }}
    <PageView />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import PageView from "@/components/PageView.vue";
import settings from "../settings";

export default defineComponent({
  name: "Game",
  components: {
    PageView,
  },
  data() {
    return {
      dev: false,
    };
  },
  mounted(): void {
    this.dev = settings.dev;
  },
});
</script>
