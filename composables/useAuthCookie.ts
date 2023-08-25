export default function () {
  return useCookie(Cookie.Auth, {
    secure: true,
    sameSite: 'strict',
  });
}
