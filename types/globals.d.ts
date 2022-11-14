type DeepPartial<T> = Partial<{ [key in keyof T]: Partial<T[key]> }>;
type DeepRequired<T> = Required<{ [key in keyof T]: Required<T[key]> }>;
