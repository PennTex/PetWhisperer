var logout = function () {
  document.cookie = "auth-session=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
  window.location.href = '/';
}