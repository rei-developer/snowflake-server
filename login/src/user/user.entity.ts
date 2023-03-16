import {
  Entity,
  Unique,
  BaseEntity,
  Column,
  PrimaryGeneratedColumn,
} from 'typeorm';

@Entity({ name: 'users' })
@Unique(['uid'])
export class User extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  uid: string;

  @Column()
  name: string;

  @Column()
  sex: number;

  @Column()
  created: Date;

  @Column()
  updated: Date;

  @Column()
  deleted: Date;
}
