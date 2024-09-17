import { AuthConfig, User } from "./types";

export class HardcodeAuth {
  private config: AuthConfig;

  constructor(config: Partial<AuthConfig>) {
    this.config = this.parseConfig(config);
  }

  async register(userData: Partial<User>): Promise<User> {
    // Pseudo code for user registration
  }

  async login(username: string, password: string): Promise<any> {}

  async refreshToken(refreshToken: string): Promise<any> {}

  parseConfig(config: Partial<AuthConfig>) {
    const defaultConfig: AuthConfig = {
      jwtSecret: "your-default-secret",
      tokenExpiry: "30m",
    };
    const cfg: AuthConfig = { ...defaultConfig, ...config };
    return cfg;
  }
}
