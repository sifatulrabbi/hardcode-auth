export type AuthConfig = {
  jwtSecret: string;
  tokenExpiry: `${number}${"m" | "h" | "d"}`;
};

export type AuthConfigOptions = {
  jwtSecret: string;
  tokenExpiry: `${number}${"m" | "h" | "d"}`;
};

export type User = {
  id: string;
  username: string;
  email: string;
  createdAt: string;
  metadata: Record<string, any>;
};

export type TokenPayload = {
  userId: string;
  email: string;
};
