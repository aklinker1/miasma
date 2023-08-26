export default function () {
  return useCookie(Cookie.Auth, {
    sameSite: 'strict',
  });
}
