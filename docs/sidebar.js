module.exports = {
  Introduction: [
    "index", "quickstart", "install",
  ],
  Concepts: [
    "concepts/index",
    "concepts/ui-user-interface",
    "concepts/identity-user-model",
    {
      label: 'Identity Credentials',
      type: "category",
      items: [
        "concepts/credentials",
        "concepts/credentials/username-email-password",
        "concepts/credentials/openid-connect-oidc-oauth2",
      ]
    },
    "concepts/selfservice-flow-completion",
    "concepts/email-sms",
    "concepts/federation",
    "concepts/security"
  ],
  "Self Service": [
    "self-service",
    {
      type: "category",
      label: "User Login and User Registration",
      items:
        [
          "self-service/flows/user-login-user-registration",
          "self-service/flows/user-login-user-registration/username-email-password",
          "self-service/flows/user-login-user-registration/openid-connect-social-sign-in-oauth2",
        ]
    },
    {
      type: "category",
      label: "User Settings",
      items:
        [
          "self-service/flows/user-settings",
          "self-service/flows/user-settings/user-profile-management",
          "self-service/flows/user-settings/change-password",
          "self-service/flows/user-settings/link-unlink-openid-connect-oauth2",
        ]
    },
    "self-service/flows/user-logout",
    "self-service/flows/password-reset-account-recovery",
    "self-service/flows/user-facing-errors",
    "self-service/flows/verify-email-account-activation",
    "self-service/flows/2fa-mfa-multi-factor-authentication",
    "self-service/hooks/index",
  ],
  Guides: [
    "guides/sign-in-with-github-google-facebook-linkedin",
    "guides/zero-trust-iap-proxy-identity-access-proxy",
    "guides/multi-tenancy-multitenant",
    "guides/high-availability-ha"
  ],
  "Reference": [
    "reference/configuration",
    "reference/json-schema-json-paths",
    "reference/html-forms",
    "reference/api"
  ],
  "Debug & Help": [
    "debug/csrf"
  ],
  SDKs: ["sdk"],
  "Further Reading": [
    "further-reading/comparison"
  ],
};
