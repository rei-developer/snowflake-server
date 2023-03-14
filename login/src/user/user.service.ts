import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { UserRepository } from './user.repository';
import { UserResponseDto } from './dto/response/user.response.dto';
import { UserModel } from './user.model';

@Injectable()
export class UserService {
  constructor(
    @InjectRepository(UserRepository)
    private readonly userRepository: UserRepository,
  ) {}

  async fetchUser(uid: string): Promise<UserResponseDto> {
    const user = await this.userRepository.readUserByUId(uid);
    console.log('fetched user : ', user);
    return new UserResponseDto(user);
  }

  async addUser(userModel: UserModel): Promise<boolean> {
    return await this.userRepository.createUser(userModel);
  }

  async patchUser(uid: string, userModel: UserModel): Promise<boolean> {
    return await this.userRepository.updateUserByUId(uid, userModel);
  }
}
