export const useAuth = () => {
  const logout = async () => {
    if (import.meta.client) {
      localStorage.removeItem("auth-token");
    }

    await navigateTo("/auth/login");
  };

  return {
    logout,
  };
};