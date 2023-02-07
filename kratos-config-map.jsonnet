function(ctx) {
  externalIdentityId: ctx.identity.id,
  email: ctx.identity.traits.email,
  firstName: std.join(" ",std.slice(std.split(ctx.identity.traits.name, " "), 0, std.length(std.split(ctx.identity.traits.name, " ")) - 1,1)),
  lastName: std.split(ctx.identity.traits.name, " ")[std.length(std.split(ctx.identity.traits.name, " ")) - 1],
  status: "ACTIVE",
  setIds: ["c9149290-2549-4bdd-a9ba-c0e02b83938a"],
  accessToken: if "access_token" in ctx.request_cookies then
    ctx.request_cookies.access_token
  else
    null,
}