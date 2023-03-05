import {
  Entity,
  Unique,
  BaseEntity,
  Column,
  PrimaryGeneratedColumn,
} from 'typeorm';
import { AuthType } from '#auth/auth.model';

@Entity({ name: 'Users' })
@Unique(['uid'])
export class User extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  type: AuthType;

  @Column()
  uid: string;

  @Column()
  email: string;

  @Column()
  created: Date;

  @Column()
  updated: Date;

  @Column()
  deleted: Date;
}
