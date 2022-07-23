<template >
  <div>
    <b-row mt="5">
      <b-col
        md="8"
        offset-md="2"
        lg="6"
        offset-lg="3"
      >
        <b-card title="登录">
          <b-form>
            <b-form-group label="手机号">
              <b-form-input
                v-model="$v.user.telephone.$model"
                type="number"
                placeholder="输入你的手机号"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('telephone')">
                请输入合法手机号
              </b-form-invalid-feedback>
              <b-form-valid-feedback :state="validateState('telephone')">
                <!-- Looks Good. -->
              </b-form-valid-feedback>
            </b-form-group>
            <b-form-group label="密码">
              <b-form-input
                v-model="$v.user.password.$model"
                type="password"
                placeholder="输入你的密码"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('password')">
                密码必须大于6位
              </b-form-invalid-feedback>
              <b-form-valid-feedback :state="validateState('password')">
                <!-- Looks Good. -->
              </b-form-valid-feedback>
            </b-form-group>
            <b-form-group>
              <b-button
                variant="outline-primary"
                block
                @click="login"
              >登录</b-button>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>
<script>
import { minLength, required } from 'vuelidate/lib/validators';
import coustomValidator from '@/helper/validator';
import { mapActions } from 'vuex';

export default {
  data() {
    return {
      user: {
        password: '',
        telephone: '',
      },
      validation: null,
    };
  },
  methods: {
    ...mapActions('userModule', { userLogin: 'login' }),
    makeToast(meg = null) {
      this.$bvToast.toast(meg, {
        title: '数据验证错误',
        variant: 'danger',
        solid: false,
      });
    },
    validateState(name) {
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    login() {
      // 验证数据
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      this.userLogin(this.user).then(() => {
        // 跳转主页
        this.$router.replace({ name: 'Home' });
      }).catch((err) => {
        console.log('err:', err.response.data.msg);
        this.makeToast(err.response.data.msg);
      });
    },
  },
  validations: {
    user: {
      telephone: {
        required,
        telephone: coustomValidator.telephoneValidator,
      },
      password: {
        minLength: minLength(6),
        required,
      },
      name: {

      },
    }, // Matches this.firstName
  },
};
</script>
<style lang="scss" scoped>
</style>
