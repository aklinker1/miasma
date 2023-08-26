import { detailedDiff } from 'deep-object-diff';
import isEqual from 'lodash.isequal';
import { MaybeRef } from 'vue';

/**
 * Handles form logic for editing a deep JSON object.
 *
 * It aliases deep feilds to simple refs that can be modified, and returns those refs and utils for
 * reseting or applying all the changes.
 */
export default function <TModel, TValues extends object>(
  prevModel: Ref<TModel>,
  getPrevValues: { [key in keyof TValues]: (model: Readonly<TModel>) => TValues[key] },
  getLatestModel: (base: TModel, values: TValues) => TModel,
  onDiscard: MaybeRef<() => void>,
) {
  const getPrevValue = <TKey extends keyof TValues>(field: TKey) =>
    clone(getPrevValues[field](toRaw(prevModel.value)));

  const fieldNames = Object.keys(getPrevValues) as Array<keyof TValues>;
  const refs = fieldNames.reduce(
    (map, field) => {
      // @ts-expect-error
      map[field] = ref(getPrevValue(field));
      return map;
    },
    {} as { [key in keyof TValues]: Ref<TValues[key]> },
  );

  const discardChanges = () => {
    fieldNames.forEach(field => {
      refs[field].value = getPrevValue(field);
    });
    unref(onDiscard)();
  };

  const latestModel = computed(() => {
    const values = Object.entries(refs).reduce((map, [key, value]) => {
      map[key as keyof TValues] = (value as Ref<any>).value;
      return map;
    }, {} as TValues);
    return getLatestModel(clone(prevModel.value), values);
  });

  const hasChanges = computed(() => !isEqual(clone(prevModel.value), clone(latestModel.value)));

  watch(
    [latestModel, hasChanges],
    ([_, newHasChanges]) => {
      if (newHasChanges)
        console.debug('Changes:', {
          diff: detailedDiff(clone(prevModel.value!), clone(latestModel.value!)),
          new: clone(latestModel.value!),
          old: clone(prevModel.value),
        });
    },
    { immediate: true },
  );

  return {
    ...refs,
    hasChanges,
    discardChanges,
    latestModel,
  };
}
