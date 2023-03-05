import { AuthType } from '#auth/auth.model';

export interface UserModel {
  type?: AuthType;
  uid?: string;
  email?: string;
}
