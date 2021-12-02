using System;
using System.Collections.Generic;
using System.IO;

namespace Aoc21
{
    public class Utils
    {
        public static IEnumerable<int> ReadIntLines(string path)
        {
            foreach (var line in File.ReadLines(path))
            {
                yield return int.Parse(line);
            }
        }
    }
}