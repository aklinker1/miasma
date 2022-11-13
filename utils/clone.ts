/**
 * Copy the object and create a mutable version of the object that can be edited.
 *
 * > `lodash.clone` as created readonly object, and cloning the input too well for updating service specs.
 */
export function clone<T>(t: T): T {
  return JSON.parse(JSON.stringify(t));
}
