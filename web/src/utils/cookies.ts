export const getCookie = (id: string) => {
  let value = document.cookie.match("(^|;)?" + id + "=([^;]*)(;|$)");
  return value ? unescape(value[2]) : null;
};
