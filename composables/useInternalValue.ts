import { Ref } from 'vue';
import isEqual from 'lodash.isequal';

export default function <K extends string, T extends Record<K, any> = Record<K, any>>(
  key: K,
  props: T,
  emit: (event: `update:${K}`, newValue: T[K]) => void,
  beforeGet?: (value: T[K]) => T[K],
  beforeEmit?: (value: T[K]) => T[K],
) {
  const v = ref(beforeGet ? beforeGet(props[key]) : props[key]) as Ref<T[K]>;

  // Update the internal value when the props change
  watch(
    () => props[key],
    newValue => {
      if (beforeGet) v.value = beforeGet(newValue);
      else v.value = newValue;
    },
  );

  // Emit the value to the parent when it changes
  watch(v, (newValue, prevValue) => {
    // For objects, don't emit if nothing has changed
    if (isEqual(newValue, prevValue)) return;

    if (beforeEmit) emit(`update:${key}`, beforeEmit(newValue));
    else emit(`update:${key}`, newValue);
  });

  return v;
}
