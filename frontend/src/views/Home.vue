<template>
  <div>
    <h1>Welcome to Home Page</h1>
    <p>This is the home page of the application.</p>
    <button @click="goToLogin">로그인</button>
    <button @click="pingApiTest">ping api test</button>
  </div>
</template>

<script>
import {useRouter} from 'vue-router';
import {axiosWithSession} from '../utils/sessionManager';

export default {
  name: 'Home',
  methods: {
    goToLogin() {
      this.$router.push('/login');
    },
    async pingApiTest() {
      try {
        const response = await axiosWithSession('/auth/echo', 'GET');
        alert(response.data);
      } catch (error) {
        console.error(error);
      }
    }
  }
};
</script>

<style scoped>
h1 {
  color: #3c3c3c;
}
</style>
