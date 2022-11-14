import { Ref } from 'vue';

type ItemOf<T> = T extends Array<infer K> ? K : never;

export default function <
  TKey extends string,
  TProps extends Record<TKey, any>,
  TItem extends ItemOf<TProps[TKey]>,
>(options: {
  key: TKey;
  props: TProps;
  emits: {
    (event: `update:${TKey}`, newValue: TProps[TKey]): void;
  };
  emptyValue: TItem;
  isEmpty: (item: TItem) => boolean;
}) {
  function trimEndOfList(prevList: TItem[] = []): TItem[] {
    const newList = [...prevList];
    // Remove all ending empty rows
    while (newList[newList.length - 1] != null && options.isEmpty(newList[newList.length - 1])) {
      newList.splice(newList.length - 1, 1);
    }
    // Add back a single empty row
    newList.push(options.emptyValue);
    return newList;
  }

  const list: Ref<TItem[]> = useInternalValue(
    options.key,
    options.props,
    options.emits,
    // Display a trimmed version of the list with an empty row, but don't emit that row
    // @ts-expect-error: types are too complex :/
    trimEndOfList,
    internalList => internalList.slice(0, internalList.length - 1),
  );

  return {
    list,
    updateItem(index: number, newValue: TItem) {
      const newList = [...list.value];
      newList[index] = newValue;
      list.value = trimEndOfList(newList);
    },
    removeItem(index: number) {
      const newList = [...list.value];
      newList.splice(index, 1);
      list.value = trimEndOfList(newList);
    },
  };
}
