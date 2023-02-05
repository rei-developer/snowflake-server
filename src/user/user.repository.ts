import { Repository } from 'typeorm';
import { CustomRepository } from '#typeorm-ex/typeorm-ex.decorator';
import { User } from './user.entity';
import { UserModel } from './user.model';

@CustomRepository(User)
export class UserRepository extends Repository<User> {
  async createUser(userModel: UserModel): Promise<boolean> {
    return !!(await this.save(this.create(userModel)));
  }

  async readUser(id: number): Promise<User> {
    return await this.findOneBy({ id });
  }

  async readUserByUId(uid: string): Promise<User> {
    return await this.findOneBy({ uid });
  }

  async updateUserByUId(uid: string, userModel: UserModel): Promise<boolean> {
    return !!(await this.update({ uid }, userModel));
  }

  async deleteUser(id: number): Promise<boolean> {
    return !!(await this.delete(id));
  }
}
