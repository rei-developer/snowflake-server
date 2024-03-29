import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { UserRepository } from './user.repository';
import { User } from './user.entity';
import { UserModel } from './user.model';

@Injectable()
export class UserService {
  constructor(
    @InjectRepository(UserRepository)
    private readonly userRepository: UserRepository,
  ) {}

  async fetchUser(uid: string): Promise<User> {
    return await this.userRepository.readUserByUId(uid);
  }

  async fetchUserByName(name: string): Promise<User> {
    return await this.userRepository.readUserByName(name);
  }

  async addUser(userModel: UserModel): Promise<number> {
    return await this.userRepository.createUser(userModel);
  }
}
