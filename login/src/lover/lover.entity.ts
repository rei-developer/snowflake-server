import { Entity, BaseEntity, Column, PrimaryGeneratedColumn } from 'typeorm';

@Entity({ name: 'lovers' })
export class Lover extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  userId: number;

  // @Column()
  // name: string;
  //
  // @Column()
  // nickname: string;
  //
  // @Column()
  // race: string;
  //
  // @Column()
  // sex: number;
  //
  // @Column()
  // age: number;
  //
  // @Column()
  // hair: string;
  //
  // @Column()
  // face: string;
  //
  // @Column()
  // eyes: string;
  //
  // @Column()
  // nose: string;
  //
  // @Column()
  // mouth: string;
  //
  // @Column()
  // ears: string;
  //
  // @Column()
  // body: string;
  //
  // @Column()
  // breast: string;
  //
  // @Column()
  // rank: string;
  //
  // @Column()
  // level: number;
  //
  // @Column()
  // exp: number;
  //
  // @Column()
  // remarks: string;
  //
  // @Column()
  // isNft: boolean;
  //
  // @Column()
  // created: Date;
  //
  // @Column()
  // updated: Date;
  //
  // @Column()
  // deleted: Date;
}
