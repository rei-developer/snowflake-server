import { Module } from '@nestjs/common';
import { TypeOrmExModule } from '#typeorm-ex/typeorm-ex.module';
import { UserRepository } from './user.repository';
import { UserService } from './user.service';
import { UserController } from './user.controller';

@Module({
  imports: [TypeOrmExModule.forCustomRepository([UserRepository])],
  exports: [UserService],
  controllers: [UserController],
  providers: [UserService],
})
export class UserModule {}
