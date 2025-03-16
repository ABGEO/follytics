import { Configuration, UserApi } from '@follytics/sdk';
import type { ConfigurationParameters } from '@follytics/sdk';

interface ApiFactoryInterface {
  updateConfiguration(config: ConfigurationParameters): void;
  updateAccessToken(accessToken?: string): void;
  getUserApi: () => UserApi;
}

class ApiFactory implements ApiFactoryInterface {
  private configuration: Configuration;
  private apiInstances: Record<string, unknown> = {};

  constructor(config: ConfigurationParameters = {}) {
    this.configuration = new Configuration(config);
  }

  public updateAccessToken(accessToken?: string): void {
    this.updateConfiguration({ ...this.configuration, apiKey: accessToken });
  }

  public updateConfiguration(config: ConfigurationParameters): void {
    this.configuration = new Configuration(config);
    this.apiInstances = {};
  }

  private getOrCreateApi<T>(key: string, factory: () => T): T {
    return (this.apiInstances[key] ??= factory()) as T;
  }

  public getUserApi(): UserApi {
    return this.getOrCreateApi('user', () => new UserApi(this.configuration));
  }
}

const createApiFactory = (config: ConfigurationParameters = {}): ApiFactory =>
  new ApiFactory(config);

export type { ApiFactoryInterface };
export { createApiFactory };
