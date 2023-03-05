import { Entity, BaseEntity, Column, PrimaryGeneratedColumn } from 'typeorm';
import { ClassType, SexType } from './word.model';

@Entity({ name: 'Words' })
export class Word extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  dicId: number;

  @Column()
  class?: ClassType;

  @Column()
  sex?: SexType;

  @Column()
  word: string;

  @Column()
  meaning: string;

  @Column()
  variation?: string;

  @Column()
  pronunciation?: string;

  @Column()
  example?: string;

  @Column()
  interpretation?: string;

  @Column()
  synonym?: string;

  @Column()
  antonym?: string;

  @Column()
  remarks?: string;

  @Column()
  isConversation: number;

  @Column()
  created: Date;

  @Column()
  updated: Date;

  @Column()
  deleted: Date;
}
