import Vue from 'vue';
import Vuex from 'vuex';
import userModule from './module/user';
import app from './module/app';

Vue.use(Vuex);

export default new Vuex.Store({
  // 开启严格模式
  strict: process.env.NODE_ENV !== 'production',
  state: {
    a: '',
  },
  mutations: {
    demo(state) {
      state.a = 'demo';
    },
  },
  actions: {
  },
  modules: {
    userModule,
    app,
  },
});
