import axios from 'axios';
import {useRouter} from 'vue-router';


export function setCookie(name: string, value: string, seconds: number) {
    let expires = "";
    if (seconds) {
        const date = new Date();
        date.setTime(date.getTime() + (seconds * 1000));
        expires = "; expires=" + date.toUTCString();
    }
    document.cookie = name + "=" + (value || "") + expires + "; path=/";
}

export function getCookie(name: string) {
    const nameEQ = name + "=";
    const ca = document.cookie.split(';');
    for (let i = 0; i < ca.length; i++) {
        let c = ca[i];
        while (c.charAt(0) == ' ') c = c.substring(1, c.length);
        if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length, c.length);
    }
    return null;
}

export function axiosWithSession(path: string, method: string, data: any) {
    const router = useRouter();
    const sessionValue = getCookie('session');
    const serverUrl = import.meta.env.VITE_SERVER_ENDPOINT;
    const fullUrl = `${serverUrl}${path}`;

    if (!sessionValue) {
        router.push('/login');
        return;
    }

    return axios({
        method: method,
        url: fullUrl,
        data: data,
        headers: {
            'Authorization': `Bearer ${sessionValue}`
        }
    }).catch(error => {
        if (error.response && error.response.status === 401) {
            router.push('/login');
        }
        throw error;
    });
}