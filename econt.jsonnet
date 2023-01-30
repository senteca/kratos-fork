local claims = std.extVar('claims');

{
  identity: {
    traits: {
      email: claims.user_name,
      user_name: claims.user_name,
      name: claims.name,
    },
  },
}