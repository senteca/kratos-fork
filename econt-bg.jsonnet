local claims = std.extVar('claims');

{
  identity: {
    traits: {
      email: claims.email,
      user_name: claims.email,
      name: claims.name,
      access_token: claims.access_token,
      expires_in: claims.expires_in,
    },
  },
}