export enum AuthType {
  GOOGLE = 'google',
  APPLE = 'apple',
}

export const CustomAuthScheme = 'snowflake';

export interface AuthModel {
  authType: AuthType;
  idToken: string;
}
