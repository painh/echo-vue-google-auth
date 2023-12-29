<template>
  <div>
    <h1>OAuth2 Callback</h1>
    <p v-if="loginStatus === 200">로그인 성공: {{ loginStatusText }}</p>

  </div>
</template>

<script lang="ts">
import {defineComponent, ref} from 'vue';
import {useRoute, useRouter} from 'vue-router';
import axios from 'axios';

const api = import.meta.env.VITE_SERVER_ENDPOINT;

export default defineComponent({
  setup() {
    const route = useRoute();
    const router = useRouter();
    const loginStatus = ref<number | null>(null);
    const loginStatusText = ref<string | null>(null);

    const code = route.query.code;

    const apiURL = `${api}/login`;
    console.log(apiURL, code);

    axios.post(apiURL, {code})
        .then(response => {
          if (response.status === 200) {
            loginStatus.value = 200;
            loginStatusText.value = response.data;
          } else {
            alert('로그인 실패');
            router.push('/login');
          }
        })
        .catch(error => {
          console.error(error);
          alert('로그인 실패');
          router.push('/login');
        });

    return {
      loginStatus,
      loginStatusText,
    };
  },
});
</script>