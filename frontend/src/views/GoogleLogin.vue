<script>
import {defineComponent} from 'vue';

import GoogleLogo from '../assets/btn_google_signin_dark_normal_web.png';
import {getGoogleUrl} from '../utils/getGoogleUrl';

const from = '/';

export default defineComponent({
  name: 'GoogleLogin',
  setup() {
    return {
      from,
      getGoogleUrl,
      GoogleLogo, // GoogleLogo를 setup 함수의 반환 객체에 추가
    };
  },
});
</script>

<template>
  <a :href="getGoogleUrl(from)" class="google-auth">
    <img :src="GoogleLogo" alt="Google Logo"/>
  </a>
</template>