<template>
  <div id="app">
      <navbar/>
      <b-container>
        <router-view/>
      </b-container>

  </div>
</template>

<script>
import Navbar from './views/layout/Navbar.vue';

export default {
  components: { Navbar },
};
</script>

<style lang="sass" scoped>

</style>
