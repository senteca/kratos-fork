local claims = std.extVar('claims');

{
  identity: {
    traits: {
      email: claims.email,
      user_name: claims.email,
      name: claims.name,
      access_token: claims.access_token,
      expires_in: claims.expires_in,
      setIds: ["c9149290-2549-4bdd-a9ba-c0e02b83938a", "6ecf0270-9350-46d8-aae5-59823a440180"],
      provider: "econtbg",
    },
  },
}