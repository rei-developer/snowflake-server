import { Repository } from 'typeorm';
import { CustomRepository } from '#typeorm-ex/typeorm-ex.decorator';
import { User } from './user.entity';
import { UserModel } from './user.model';

@CustomRepository(User)
export class UserRepository extends Repository<User> {
  async createUser(userModel: UserModel): Promise<number | boolean> {
    try {
      const user = this.create(userModel);
      await this.save(user);
      return user.id;
    } catch (err) {
      return false;
    }
  }

  async readUserByUId(uid: string): Promise<User> {
    return await this.findOneBy({ uid });
  }
}
