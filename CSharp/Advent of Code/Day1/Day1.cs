using System;
using System.Collections.Generic;
using System.Linq;

// HERE LIES MADNESS, YOU HAVE BEEN WARNED
namespace Aoc21
{
    public class Day1
    {
        public static int Part1(string path)
        {
            var prev = 0;
            var increaseCount = 0;
            foreach (var current in Utils.ReadIntLines(path))
            {
                if (prev != 0 && current - prev > 0)
                {
                    increaseCount++;
                }

                prev = current;
            }

            return increaseCount;
        }

        public static int Part2(string path)
        {
            var IntervalA = new Interval(0);
            var IntervalB = new Interval(-1);
            var IntervalC = new Interval(-2);

            var increaseCount = 0;
            var counter = -1;
            var intervalSums = new Dictionary<string, int>
            {
                {"A", 0},
                {"B", 0},
                {"C", 0}
            };
            foreach (var line in Utils.ReadIntLines(path))
            {
                counter++;
                // populate windows
                IntervalA.Push(counter, line);
                IntervalB.Push(counter, line);
                IntervalC.Push(counter, line);

                if (counter < 2)
                {
                    continue;
                }

                var iteration = counter % 3;
                switch (iteration)
                {
                    case 0:
                        // interval B has just filled
                        var sumB = IntervalB.Window.Sum();
                        intervalSums.TryAdd("B", sumB);
                        intervalSums.TryGetValue("A", out var sumA);
                        if (sumB > sumA)
                        {
                            increaseCount++;
                        }

                        break;
                    case 1:
                        // interval C has just filled, so let's compare it to interval B
                        var sumC = IntervalC.Window.Sum();
                        intervalSums.TryAdd("C", sumC);
                        intervalSums.TryGetValue("B", out sumB);
                        if (sumC > sumB)
                        {
                            increaseCount++;
                        }

                        break;
                    case 2:
                        // interval A has just filled, so let's compare it to interval C
                        sumA = IntervalC.Window.Sum();
                        intervalSums.TryAdd("A", sumA);
                        intervalSums.TryGetValue("C", out sumC);
                        if (sumC == 0)
                        {
                            break;
                        }

                        if (sumA > sumC)
                        {
                            increaseCount++;
                        }

                        break;
                }

                counter++;
            }

            return increaseCount;
        }
    }

    class Interval
    {
        public int[] Window;
        private int Offset;

        public Interval(int offset)
        {
            Window = new int[3];
            Offset = offset;
        }

        public void Push(int index, int value)
        {
            var calculatedIndex = Math.Abs((index + Offset) % 3);
            Window[calculatedIndex] = value;
        }
    }
}