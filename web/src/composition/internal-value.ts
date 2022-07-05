import { Ref } from "vue";

export function useInternalValue<
  K extends string,
  T extends Record<K, any> = Record<K, any>
>(key: K, props: T, emit: (event: `update:${K}`, newValue: T[K]) => void) {
  const v = ref(props[key]) as Ref<T[K]>;

  // Update the internal value when the props change
  watch(
    () => props[key],
    (newValue) => {
      v.value = newValue;
    }
  );
  // Emit the value to the parent when it changes
  watch(v, (newValue) => {
    emit(`update:${key}`, newValue);
  });

  return v;
}
