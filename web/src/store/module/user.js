import storageService from '@/service/storageService';
import userService from '@/service/userService';

const userModule = {
  namespaced: true,
  state: {
    token: storageService.get(storageService.USER_TOKEN),
    userInfo: JSON.parse(storageService.get(storageService.USER_INFO)),
    demo: '',
  },
  mutations: {
    SET_TOKEN(state, token) {
      // 更新本地缓存
      storageService.set(storageService.USER_TOKEN, token);
      state.token = token;
    },
    SET_USERINFO(state, userInfo) {
      // 更新本地缓存
      storageService.set(storageService.USER_INFO, JSON.stringify(userInfo));
      state.userInfo = userInfo;
    },

  },
  actions: {
    register(context, { username, telephone, password }) {
      return new Promise((resolve, reject) => {
        userService.register({ username, telephone, password }).then((res) => {
          // 保存token
          // 用 vuex
          context.commit('SET_TOKEN', res.data.data.token);
          // storageService.set(storageService.USER_TOKEN, res.data.data.token);
          return userService.info();
        }).then((response) => {
          // 保存用户信息 JSON序列化
          context.commit('SET_USERINFO', response.data.data.data.user);
          resolve(response);
        }).catch((err) => {
          reject(err);
        });
      });
    },
    login(context, { telephone, password }) {
      return new Promise((resolve, reject) => {
        userService.login({ telephone, password }).then((res) => {
          context.commit('SET_TOKEN', res.data.data.token);
          return userService.info();
        }).then((response) => {
          context.commit('SET_USERINFO', response.data.data.data.user);
          resolve(response);
        }).catch((err) => {
          reject(err);
        });
      });
    },
    logout({ commit }) {
      commit('SET_TOKEN', '');
      storageService.set(storageService.USER_TOKEN, '');

      commit('SET_USERINFO', null);
      storageService.set(storageService.USER_INFO, null);
      window.location.reload();
    },
  },
};
export default userModule;
