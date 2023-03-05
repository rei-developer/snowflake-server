import { Module } from '@nestjs/common';
import { FirebaseAuthStrategy } from '#firebase/firebase-auth.strategy';

@Module({
  exports: [FirebaseAuthStrategy],
  providers: [FirebaseAuthStrategy],
})
export class FirebaseAuthModule {}
