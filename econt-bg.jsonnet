local claims = std.extVar('claims');

{
  identity: {
    traits: {
      email: claims.email,
      user_name: claims.email,
      name: claims.name,
    },
  },
}