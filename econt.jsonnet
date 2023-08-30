local claims = std.extVar('claims');

{
  identity: {
    traits: {
      email: claims.user_name,
      user_name: claims.user_name,
      name: claims.name,
      access_token: claims.access_token,
      expires_in: claims.expires_in,
      provider: "econt",
    },
  },
}