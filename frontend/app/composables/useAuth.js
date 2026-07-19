import { computed } from "vue";

export const useAuth = () => {
  const token = useState("auth-token", () => {
    if (import.meta.client) {
      return localStorage.getItem("auth-token");
    }

    return null;
  });

  const loggedIn = computed(() => !!token.value);

  const login = (newToken) => {
    token.value = newToken;

    if (import.meta.client) {
      localStorage.setItem("auth-token", newToken);
    }
  };

  const logout = async () => {
    token.value = null;

    if (import.meta.client) {
      localStorage.removeItem("auth-token");
    }

    await navigateTo("/auth/login");
  };

  return {
    token,
    loggedIn,
    login,
    logout,
  };
};
