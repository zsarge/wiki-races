<template>
  <div class="game">
    This is the view of the page:
    <PageView />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import PageView from "@/components/PageView.vue";

export default defineComponent({
  name: "Game",
  components: {
    PageView,
  },
});
</script>
