import { detailedDiff } from 'deep-object-diff';
import isEqual from 'lodash.isequal';
import { MaybeRef } from 'vue';

/**
 * Handles form logic for editing a deep JSON object.
 */
export default function <TModel, TFields extends string>(
  prevModel: Ref<TModel>,
  getPrevValues: Record<TFields, (model: Readonly<TModel>) => any>,
  getLatestModel: (base: TModel, values: Record<TFields, any>) => TModel,
  onDiscard: MaybeRef<() => void>,
) {
  const getPrevValue = (field: TFields) => clone(getPrevValues[field](toRaw(prevModel.value)));

  const fieldNames = Object.keys(getPrevValues) as Array<TFields>;
  const refs = fieldNames.reduce((map, field) => {
    map[field] = ref(getPrevValue(field));
    return map;
  }, {} as Record<TFields, Ref<any>>);

  const discardChanges = () => {
    fieldNames.forEach(field => {
      refs[field].value = getPrevValue(field);
    });
    unref(onDiscard)();
  };

  const latestModel = computed(() => {
    const values = Object.entries(refs).reduce((map, [key, value]) => {
      map[key as TFields] = (value as Ref<any>).value;
      return map;
    }, {} as Record<TFields, any>);
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
