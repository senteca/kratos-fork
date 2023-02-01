function(ctx) {
  externalIdentityId: ctx.identity.id,
  email: ctx.identity.traits.email,
  status: "ACTIVE",
  setIds: ["c9149290-2549-4bdd-a9ba-c0e02b83938a"],
  accessToken: if "access_token" in ctx.request_cookies then
    ctx.request_cookies.access_token
  else
    null,
}