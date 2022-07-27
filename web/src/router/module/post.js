const postRoutes = [
  {
    path: '/post',
    name: 'post',
    meta: {
      auth: true,
    },
    component: () => import('@/views/post/PostShow.vue'),
    children: [
      {
        path: 'create',
        name: 'create',
        meta: {
          auth: true,
        },
        component: () => import('@/views/post/PostCreate.vue'),
      },
      {
        path: 'list',
        name: 'list',
        meta: {
          auth: true,
        },
        component: () => import('@/views/post/PostList.vue'),
      },
    //   {
    //     path: 'show',
    //     name: 'show',
    //     component: () => import('@/views/post/PostShow.vue'),
    //   },
    ],
  },

];
export default postRoutes;
