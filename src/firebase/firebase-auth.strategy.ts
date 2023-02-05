import { Injectable, UnauthorizedException } from '@nestjs/common';
import { PassportStrategy } from '@nestjs/passport';
import { Strategy, ExtractJwt } from 'passport-firebase-jwt';
import { AuthType } from '#auth/auth.model';
import * as firebase from 'firebase-admin';
import * as serviceAccountConfig from '#configs/service_account.json';

const serviceAccount = {
  type: serviceAccountConfig.type,
  projectId: serviceAccountConfig.project_id,
  privateKeyId: serviceAccountConfig.private_key_id,
  privateKey: serviceAccountConfig.private_key,
  clientEmail: serviceAccountConfig.client_email,
  clientId: serviceAccountConfig.client_id,
  authUri: serviceAccountConfig.auth_uri,
  tokenUri: serviceAccountConfig.token_uri,
  authProviderX509CertUrl: serviceAccountConfig.auth_provider_x509_cert_url,
  clientC509CertUrl: serviceAccountConfig.client_x509_cert_url,
};

@Injectable()
export class FirebaseAuthStrategy extends PassportStrategy(
  Strategy,
  'firebase-auth',
) {
  private defaultApp: any;

  constructor() {
    super({
      jwtFromRequest: ExtractJwt.fromExtractors(
        Object.keys(AuthType).map((item) =>
          ExtractJwt.fromAuthHeaderWithScheme(item),
        ),
      ),
    });
    this.defaultApp = firebase.initializeApp({
      credential: firebase.credential.cert(serviceAccount),
    });
  }

  async validate(idToken: string) {
    const firebaseUser: any = await this.defaultApp
      .auth()
      .verifyIdToken(idToken, true)
      .catch((err) => {
        console.log(err);
        throw new UnauthorizedException(err.message);
      });
    if (!firebaseUser) {
      throw new UnauthorizedException();
    }
    return firebaseUser;
  }
}
