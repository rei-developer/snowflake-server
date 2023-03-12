export enum AuthType {
  GOOGLE = 'google',
  APPLE = 'apple',
  FACEBOOK = 'facebook',
  TWITTER = 'twitter',
}

export const CustomAuthScheme = 'local';

export interface AuthModel {
  authType: AuthType;
  idToken: string;
}
