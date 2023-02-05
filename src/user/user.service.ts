import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { UserRepository } from './user.repository';
import { UserModel } from './user.model';

@Injectable()
export class UserService {
  constructor(
    @InjectRepository(UserRepository)
    private readonly userRepository: UserRepository,
  ) {}

  async addUser(userModel: UserModel): Promise<boolean> {
    return await this.userRepository.createUser(userModel);
  }

  async patchUser(uid: string, userModel: UserModel): Promise<boolean> {
    return await this.userRepository.updateUserByUId(uid, userModel);
  }
}
