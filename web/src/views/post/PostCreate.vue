<template lang="">
  <div>
    <b-form>
      <b-form-group label="标题"  label-cols-lg="1">
        <b-input-group>
          <b-form-input id="form-image" :disabled="busy"
          accept="image/*" v-model="post.title"></b-form-input>
        </b-input-group>
      </b-form-group>
      <b-form-group label="分类"  label-cols-lg="1">
             <b-form-input list="my-list-id" v-model="post.category"></b-form-input>
             <datalist id="my-list-id">
              <option>Manual Option</option>
            <option v-bind:key="category"
            v-for="category in categoryOptions">{{category}}</option>
            </datalist>
      </b-form-group>
      <b-form-group label="封面图" label-for="form-image" label-cols-lg="1">
        <b-input-group>
          <b-form-file accept="image/*" placeholder="Choose a file or drop it here..."
           v-model="post.coverImg"></b-form-file>
        </b-input-group>
      </b-form-group>
      <b-form-group label="内容">
              <b-form-input
                v-model="post.content"
                type="text"
              ></b-form-input>
      </b-form-group>
    </b-form>
  </div>
</template>
<script>
export default {
  data() {
    return {
      busy: false,
      post: {
        title: '',
        coverImg: null,
        content: '',
        category: '',
        categoryId: '',
      },
      categoryOptions: ['Small', 'Medium', 'Large', 'Extra Large'],
    };
  },
};
</script>
<style lang="sass" scoped>

</style>
