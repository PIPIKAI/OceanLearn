<!-- eslint-disable max-len -->
<template lang="">
<div>

  <b-navbar toggleable="lg" type="dark" variant="info">
    <b-container>
    <b-navbar-brand @click="$router.push({name:'Home'},onComplete => {},
onAbort => {})">Logo</b-navbar-brand>

    <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

    <b-collapse id="nav-collapse" is-nav>

      <!-- Right aligned nav items -->
      <b-navbar-nav class="ml-auto">
        <!-- <b-nav-form>
          <b-form-input size="sm" class="mr-sm-2" placeholder="Search"></b-form-input>
          <b-button size="sm" class="my-2 my-sm-0" type="submit">Search</b-button>
        </b-nav-form> -->

        <b-nav-item-dropdown v-if="userInfo" right>
          <template #button-content>
            <em>{{userInfo.username}}</em>
          </template>
          <b-dropdown-item @click="$router.replace({name:'profile'})">个人主页</b-dropdown-item>
          <b-dropdown-item  @click="logout">登出</b-dropdown-item>
        </b-nav-item-dropdown>
        <div >
        <b-nav-item v-if="$route.name!='login'" @click="$router.replace({name:'login'})">登陆</b-nav-item>
        <b-nav-item right v-if="$route.name!='register'" @click="$router.replace({name:'register'})">注册</b-nav-item>
        </div>

      </b-navbar-nav>
    </b-collapse>
    </b-container>
  </b-navbar>

</div>
</template>
<script>
import { mapActions, mapState } from 'vuex';

export default {
  computed: mapState({
    // userInfo: (state) => (state.userModule.userInfo === '' ? {} : state.userModule.userInfo),
    userInfo: (state) => state.userModule.userInfo,
  }),
  methods: {
    ...mapActions('userModule', { userLogout: 'logout' }),
    logout() {
      this.userLogout();
    },
  },

};
</script>
<style lang="sass" scoped>
</style>
